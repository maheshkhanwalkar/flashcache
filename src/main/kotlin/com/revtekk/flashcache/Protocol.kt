package com.revtekk.flashcache

import java.nio.ByteBuffer
import java.nio.channels.SocketChannel
import java.util.concurrent.atomic.AtomicBoolean

/**
 * Read a command from the client channel
 *
 * For simplicity, this method returns null if it encounters any kind of error,
 * whether it be I/O related or bad format
 */
internal fun readCommand(client: SocketChannel, quit: AtomicBoolean): Command? {
    val szBytes = forceReadOrNull(client, quit, 4) ?: return null
    val size = szBytes.int

    if(size < 0) {
        return null
    }

    val typeBytes = forceReadOrNull(client, quit, 1) ?: return null
    val type = CommandType.fromRaw(typeBytes.get()) ?: return null

    val keySzBytes = forceReadOrNull(client, quit, 4) ?: return null
    val keySz = keySzBytes.int

    if(keySz <= 0) {
        return null
    }

    val keyBytes = forceReadOrNull(client, quit, keySz) ?: return null
    val key = String(keyBytes.array())

    val value = if(size == 0) {
        null
    } else {
        val raw = forceReadOrNull(client, quit, size) ?: return null
        raw.array().toList()
    }

    return Command(type, key, value)
}

/**
 * Read the entire amount specified and return the byte buffer
 *
 * If there is an exception or no more data available, then the method will return null, since it
 * could not completely fulfill the request
 */
private fun forceReadOrNull(client: SocketChannel, quit: AtomicBoolean, amount: Int): ByteBuffer? {
    val buffer = ByteBuffer.allocateDirect(amount)

    try {
        var total = 0
        var curr = client.read(buffer)

        if(curr == -1)
            return null

        total += curr

        while(total < amount) {
            if(quit.get())
                return null

            curr = client.read(buffer)

            if(curr == -1)
                return null

            total += curr
        }

        buffer.flip()
        return buffer
    }
    catch (ex: Exception) {
        return null
    }
}
