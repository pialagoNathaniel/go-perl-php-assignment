# Pialago, Nathaniel Christian M.  BSCS601
my $test = 4;
print $test;
$test = "test";
print $test;
print "Good Morning! Did you sleep well? [y/n] ";
chomp(my $choice = <STDIN>);

if ($choice eq "y") {
    print "Great! Now do 30 minutes of exercise and relax!\n";
} else {
    print "That's okay! Do 15 minutes of meditation!\n";
}

print "Have a breakfast!\n";
print "Did you take a shower in the morning? [y/n] ";
chomp($choice = <STDIN>);

if ($choice eq "y") {
    print "Cool!";
} else {
    print "Alright! Take a shower!";
}

print "And get dressed!\n";
print "Do you feel ready to start the day? [y/n] ";
chomp($choice = <STDIN>);

if ($choice eq "y") {
    print "That's great! You should make a to-do list!\n";
} else {
    print "That's okay! Take 15 minutes for self-care or self-love!\n";
}

print "Are you feeling better now? [y/n]";
chomp($choice = <STDIN>);

if ($choice eq "y") {
    print "Awesome! You are now ready for a strong day!";
} else {
    print "I see... Take a test!";
}