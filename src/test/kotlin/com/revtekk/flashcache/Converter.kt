package com.revtekk.flashcache

import java.nio.ByteBuffer

fun fromIntToBytes(num: Int): List<Byte> {
    val buffer = ByteBuffer.allocate(4)
    buffer.putInt(num)

    return buffer.array().toList()
}

fun fromBytesToInt(bytes: List<Byte>): Int {
    val buffer = ByteBuffer.wrap(bytes.toByteArray())
    return buffer.int
}

fun fromStringToBytes(str: String): List<Byte> =
    str.toByteArray().toList()

fun fromBytesToString(bytes: List<Byte>): String =
    String(bytes.toByteArray())
