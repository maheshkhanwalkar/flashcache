package com.revtekk.flashcache

enum class CommandType(val value: Byte) {
    GET(0), PUT(1);

    companion object {
        fun fromRaw(value: Byte): CommandType? = values().firstOrNull { type -> type.value == value }
    }
}

data class Command(val type: CommandType, val key: String, val value: List<Byte>?)
