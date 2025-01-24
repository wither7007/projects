
import java.io.File

fun main() {
    val directoryPath = "exampleDir" // Replace with your directory path
    val directory = File(directoryPath)

    val files = directory.listFiles()?.filter { it.isFile }
    files?.forEach { file ->
        println(file.name)
    } ?: println("No files found or directory does not exist.")
}
