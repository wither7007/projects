package com.baeldung.run

class RunClass {
    fun printInsideClass() {
        println("Running inside the RunClass")
    }
}

fun main(args: Array<String>) {
    println("Running the main function")
    RunClass().printInsideClass()
}
