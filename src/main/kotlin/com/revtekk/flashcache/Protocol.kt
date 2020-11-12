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
        raw.array()
    }

    return Command(type, key, value)
}

/**
 * Write the command to the server channel
 */
internal fun writeCommand(server: SocketChannel, cmd: Command) {
    val size = if(cmd.value == null) 0 else cmd.value.size
    val type = cmd.type.value
    val keySize = cmd.key.length

    val buffer = ByteBuffer.allocateDirect(4 + 1 + 4 + keySize + size)

    buffer.putInt(size)
    buffer.put(type)
    buffer.putInt(keySize)
    buffer.put(cmd.key.toByteArray())

    if(cmd.value != null)
        buffer.put(cmd.value)

    buffer.flip()
    server.write(buffer)
}

/**
 * Read a response from the server channel
 *
 * For simplicity, this method returns null if it encounters any kind of error,
 * whether it be I/O related or bad format
 */
internal fun readResponse(client: SocketChannel): Response? {
    val fake = AtomicBoolean()

    val szBytes = forceReadOrNull(client, fake, 4) ?: return null
    val size = szBytes.int

    if(size < 0) {
        return null
    }

    val typeBytes = forceReadOrNull(client, fake, 1) ?: return null
    val type = ResponseType.fromRaw(typeBytes.get()) ?: return null

    val data = if(size == 0) {
        null
    } else {
        val raw = forceReadOrNull(client, fake, size) ?: return null
        raw.array()
    }

    return Response(type, data)
}

/**
 * Write the response to the client channel
 */
internal fun writeResponse(client: SocketChannel, resp: Response) {
    val size = if(resp.data == null) 0 else resp.data.size
    val type = resp.type.value

    val buffer = ByteBuffer.allocateDirect(4 + type + size)

    buffer.putInt(size)
    buffer.put(type)

    if(resp.data != null)
        buffer.put(resp.data)

    buffer.flip()
    client.write(buffer)
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
