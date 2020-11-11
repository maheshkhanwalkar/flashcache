package com.revtekk.flashcache

data class Response(val type: ResponseType, val data: List<Byte>?)

enum class ResponseType(val value: Byte) {
    OK(0), DATA(1), ERR(2);

    companion object {
        fun fromRaw(value: Byte): ResponseType? = values().firstOrNull { tp -> tp.value == value }
    }
}
