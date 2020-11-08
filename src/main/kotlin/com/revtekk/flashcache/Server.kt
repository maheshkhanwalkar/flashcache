package com.revtekk.flashcache

import java.net.InetSocketAddress
import java.nio.channels.SelectionKey
import java.nio.channels.SelectionKey.*
import java.nio.channels.Selector
import java.nio.channels.ServerSocketChannel
import java.nio.channels.SocketChannel
import java.util.concurrent.atomic.AtomicBoolean

class Server(private val port: Int) {
    private lateinit var socket: ServerSocketChannel
    private lateinit var selector: Selector

    private val quit = AtomicBoolean(false)

    fun start() {
        socket = ServerSocketChannel.open().bind(InetSocketAddress(port))
        selector = Selector.open()

        // Setup the newly created server channel
        with(socket) {
            configureBlocking(false)
            register(selector, OP_ACCEPT)
        }

        // Enter the event loop
        while(!quit.get()) {
            with(selector) {
                select()
                val keys = selectedKeys()

                for(key in keys) {
                    if(key.isAcceptable)
                        acceptClient(key)

                    if(key.isReadable)
                        processClient(key)
                }
            }
        }
    }

    private fun acceptClient(key: SelectionKey) {
        val client = socket.accept()

        with(client) {
            configureBlocking(false)
            register(selector, OP_READ)
        }
    }

    private fun processClient(key: SelectionKey) {
        val client = key.channel() as SocketChannel
        val cmd = readCommand(client, quit)

        // Close the client on error
        if(cmd == null) {
            key.cancel()
            client.close()
            return
        }

        TODO("execute the command")
    }
}
