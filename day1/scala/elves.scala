import scala.io.Source

// Define a data structure for an Elf.
case class Elf(number: Int, cals: Int)

object elves extends App {
  var elves: List[Elf] = List.empty

  // Read the input data into a list, with a blank line separating the entries
  // between the calorie records for each elf.
  val inputData: Array[String] = Source.fromFile("../input.dat").mkString.split("\n\n")

  // Calculate the total calories for each elf.
  for ((entries, i) <- inputData.zipWithIndex) {
    // New elf.
    var elf: Elf = Elf(i + 1, 0)
    // Calculate all this elf's calories.
    for (entry <- entries.split("\n")) {
      val cals: Int = entry.toInt
      elf = elf.copy(cals = elf.cals + cals)
    }
    // Add this elf to the list.
    elves = elf :: elves
  }

  // Sort the elf list descending by total calories.
  elves = elves.sortBy(_.cals)(Ordering[Int].reverse)

  // Part 1.
  val topElf: Elf = elves.head
  println(s"Part1: Elf ${topElf.number} has ${topElf.cals} calories")

  // Part 2.
  val total: Int = elves.take(3).map(_.cals).sum
  println(s"Part2: Top three elves have $total calories")
}

