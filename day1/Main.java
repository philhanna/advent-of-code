import java.io.File;
import java.io.IOException;
import java.util.List;
import java.util.ArrayList;
import java.util.Scanner;

// Define a data structure for an Elf.
class Elf {
    int number;
    int cals;

    public Elf(int number, int cals) {
        this.number = number;
        this.cals = cals;
    }
}

public class Main {
    public static void main(String[] args) throws IOException {
        // Make a list to keep track of elf numbers and their total calories.
        List<Elf> elves = new ArrayList<Elf>();

        // Read the input data into a list, with a blank line separating the entries
        // between the calorie records for each elf.
        Scanner scanner = new Scanner(new File("input.dat"));
        scanner.useDelimiter("\n\n");
        while (scanner.hasNext()) {
            String entries = scanner.next();
            int i = elves.size() + 1;
            Elf elf = new Elf(i, 0);
            // Calculate all this elf's calories.
            Scanner entryScanner = new Scanner(entries);
            while (entryScanner.hasNext()) {
                int cals = entryScanner.nextInt();
                elf.cals += cals;
            }
            entryScanner.close();
            // Add this elf to the list.
            elves.add(elf);
        }
        scanner.close();

        // Sort the elf list descending by total calories.
        elves.sort((e1, e2) -> Integer.compare(e2.cals, e1.cals));

        // Part 1
        Elf topElf = elves.get(0);
        System.out.printf("Part1: Elf %d has %d calories\n", topElf.number, topElf.cals);

        // Part 2
        int total = elves.get(0).cals + elves.get(1).cals + elves.get(2).cals;
        System.out.printf("Part2: Top three elves have %d calories\n", total);
    }
}

