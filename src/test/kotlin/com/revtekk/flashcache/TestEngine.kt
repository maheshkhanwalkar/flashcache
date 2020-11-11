package com.revtekk.flashcache

import org.junit.Test
import org.junit.Assert.*

class TestEngine {
    private val engine = Engine()

    @Test
    fun testGetPut()
    {
        val putRes = engine.execute(Command(CommandType.PUT, "my-key", "my-value".toByteArray().toList()))
        assertEquals(putRes.type, ResponseType.OK)

        val getRes = engine.execute(Command(CommandType.GET, "my-key", null))

        assertEquals(getRes.type, ResponseType.DATA)
        assertNotNull(getRes.data)
        assertEquals(fromBytesToString(getRes.data!!), "my-value")
    }

    @Test
    fun testMissingKey()
    {
        val res = engine.execute(Command(CommandType.GET, "my-key", null))
        assertEquals(res.type, ResponseType.ERR)
    }
}
