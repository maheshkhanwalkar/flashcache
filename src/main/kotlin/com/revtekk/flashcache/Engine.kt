package com.revtekk.flashcache

class Engine {
    private val map = mutableMapOf<String, List<Byte>>()

    fun execute(cmd: Command): Response {
        return when(cmd.type) {
            CommandType.GET -> Response(ResponseType.DATA, map[cmd.key])
            CommandType.PUT -> {
                if(cmd.value != null) {
                    map[cmd.key] = cmd.value
                }

                Response(ResponseType.OK, null)
            }
        }
    }
}
