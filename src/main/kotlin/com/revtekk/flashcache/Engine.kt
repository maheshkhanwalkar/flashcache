package com.revtekk.flashcache

class Engine {
    private val map = mutableMapOf<String, ByteArray>()

    fun execute(cmd: Command): Response {
        return when(cmd.type) {
            CommandType.GET -> executeGet(cmd.key)
            CommandType.PUT -> executePut(cmd.key, cmd.value)
            CommandType.CONTAINS -> executeContains(cmd.key)
            CommandType.DELETE -> executeDelete(cmd.key)
        }
    }

    private fun executeGet(key: String): Response =
        if (key !in map)
            Response(ResponseType.ERR, "key does not exist".toByteArray())
        else
            Response(ResponseType.DATA, map[key])

    private fun executePut(key: String, value: ByteArray?): Response =
        if(value == null)
            Response(ResponseType.ERR, "no value provided".toByteArray())
        else {
            map[key] = value
            Response(ResponseType.OK, null)
        }

    private fun executeContains(key: String): Response {
        val type = if (key in map) ResponseType.TRUE else ResponseType.FALSE
        return Response(type, null)
    }

    private fun executeDelete(key: String): Response {
        val value = map.remove(key) ?: return Response(ResponseType.ERR, "key does not exist".toByteArray())
        return Response(ResponseType.DATA, value)
    }
}
