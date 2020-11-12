package com.revtekk.flashcache

import java.nio.ByteBuffer

fun fromIntToBytes(num: Int): ByteArray {
    val buffer = ByteBuffer.allocate(4)
    buffer.putInt(num)

    return buffer.array()
}

fun fromBytesToInt(bytes: ByteArray): Int {
    val buffer = ByteBuffer.wrap(bytes)
    return buffer.int
}

fun fromStringToBytes(str: String): ByteArray =
    str.toByteArray()

fun fromBytesToString(bytes: ByteArray): String =
    String(bytes)
