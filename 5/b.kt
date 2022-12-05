// The way kotlin does arrays is so dumb
// This problem took me hours ages I couldn't quickly iterate because
// Kotlin's compile times are huge

// This was my first venture into regex

import java.io.File
import java.io.BufferedReader
import java.util.Stack

fun main() {
  val bufferedReader: BufferedReader = File("input.txt").bufferedReader()    
  val inputString = bufferedReader.use { it.readText() }
  bufferedReader.close()
  
  val halves = inputString.split("\n\n")
  var crates: Array<Stack<String>> = arrayOf()

  val lines = halves[0].split("\n")
  "[0-9]".toRegex().findAll(lines[lines.size - 1]).forEach {
    crates = crates.plus(Stack<String>())
  }

  for (i in lines.size - 2 downTo 0 step 1) {
    "[A-Z]| {4}".toRegex().findAll(lines[i]).forEachIndexed { index, regexMatch ->
      if (regexMatch.value.matches("[A-Z]".toRegex())) {
        crates[index].push(regexMatch.value)
      }
    }
  }

  halves[1].split("\n").forEach {
    val words = it.split(" ")
    val numberOfCratesToMove = words[1].toInt()
    val stack = crates[words[3].toInt() - 1]
    val stackSize = stack.count()
    
    for (i in 1 .. numberOfCratesToMove) {
      val popped = stack.removeAt(stackSize - numberOfCratesToMove)
      crates[words[5].toInt() - 1].push(popped)
    }
  }

  crates.forEach {
    print(it.peek())
  }
  
  print("\n")
}