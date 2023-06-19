# Define a data structure for an Elf.
package Elf;
sub new {
    my $class = shift;
    my $self = {
        number => shift,
        cals => shift,
    };
    bless $self, $class;
    return $self;
}

# Make an array to keep track of elf numbers and their total calories.
my @elves;

# Read the input data into an array, with a blank line separating the entries
# between the calorie records for each elf.
open(my $fh, "<", "input.dat") or die "Cannot open file: $!";
my $inputData = do { local $/; <$fh> };
close($fh);
$inputData =~ s/\r//g;
my @inputData = split(/\n\n/, $inputData);

# Calculate the total calories for each elf.
for my $i (0..$#inputData) {
    # New elf.
    my $elf = Elf->new($i + 1, 0);
    # Calculate all this elf's calories.
    my @entries = split(/\n/, $inputData[$i]);
    for my $entry (@entries) {
        my $cals = int($entry);
        $elf->{cals} += $cals;
    }
    # Add this elf to the array.
    push(@elves, $elf);
}

# Sort the elf array descending by total calories.
@elves = sort { $b->{cals} <=> $a->{cals} } @elves;

# Part 1.
my $topElf = $elves[0];
printf("Part1: Elf %d has %d calories\n", $topElf->{number}, $topElf->{cals});

# Part 2.
my $total = 0;
for my $i (0..2) {
    $total += $elves[$i]->{cals};
}
printf("Part2: Top three elves have %d calories\n", $total);
