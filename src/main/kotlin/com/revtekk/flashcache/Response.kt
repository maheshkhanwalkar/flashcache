package com.revtekk.flashcache

class Response(val type: ResponseType, val data: ByteArray?)

enum class ResponseType(val value: Byte) {
    OK(0), DATA(1), ERR(2);

    companion object {
        fun fromRaw(value: Byte): ResponseType? = values().firstOrNull { tp -> tp.value == value }
    }
}
