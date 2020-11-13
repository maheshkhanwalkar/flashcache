package com.revtekk.flashcache

import org.junit.Assert.assertEquals
import org.junit.Assert.assertNotNull
import org.junit.Before
import org.junit.Test

class TestEngine {
    private lateinit var engine: Engine

    @Before
    fun initEngine() {
        engine = Engine()
    }

    @Test
    fun testGetPut()
    {
        val putRes = engine.execute(Command(CommandType.PUT, "my-key", "my-value".toByteArray()))
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

    @Test
    fun testReplacement()
    {
        engine.execute(Command(CommandType.PUT, "key", "value0".toByteArray()))
        engine.execute(Command(CommandType.PUT, "key", "value1".toByteArray()))

        val get = engine.execute(Command(CommandType.GET, "key", null))

        assertEquals(get.type, ResponseType.DATA)
        assertNotNull(get.data)
        assertEquals(fromBytesToString(get.data!!), "value1")
    }

    @Test
    fun testContains()
    {
        var res = engine.execute(Command(CommandType.CONTAINS, "key", null))
        assertEquals(res.type, ResponseType.FALSE)

        engine.execute(Command(CommandType.PUT, "key", "value".toByteArray()))
        res = engine.execute(Command(CommandType.CONTAINS, "key", null))
        assertEquals(res.type, ResponseType.TRUE)
    }

    @Test
    fun testDelete()
    {
        engine.execute(Command(CommandType.PUT, "key", "value".toByteArray()))
        var res = engine.execute(Command(CommandType.DELETE, "key", null))

        assertEquals(res.type, ResponseType.DATA)
        assertEquals(fromBytesToString(res.data!!), "value")

        res = engine.execute(Command(CommandType.CONTAINS, "key", null))
        assertEquals(res.type, ResponseType.FALSE)
    }
}
