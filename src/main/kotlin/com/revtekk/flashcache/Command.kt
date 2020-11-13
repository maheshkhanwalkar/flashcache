package com.revtekk.flashcache

class Command(val type: CommandType, val key: String, val value: ByteArray?)

enum class CommandType(val value: Byte) {
    GET(0), PUT(1), CONTAINS(2), DELETE(3);

    companion object {
        fun fromRaw(value: Byte): CommandType? = values().firstOrNull { type -> type.value == value }
    }
}
