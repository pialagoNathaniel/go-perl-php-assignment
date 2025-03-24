<?php
// Pialago, Nathaniel Christian M.  BSCS601
function ask($question) {
    echo $question . " [y/n] ";
    $answer = trim(fgets(STDIN));
    return strtolower($answer);
}

$choice = ask("Good Morning! Did you sleep well?");
if ($choice === "y") {
    echo "That's good! Now do 30 minutes of exercise and then relax!\n";
} else {
    echo "That's alright! Do 15 minutes of meditation!\n";
}

echo "Now have a breakfast!\n";
$choice = ask("Did you take a shower in the morning?");
if ($choice === "y") {
    echo "Nice! ";
} else {
    echo "Alright! Take a shower! ";
}

echo "And get dressed!\n";
$choice = ask("Do you feel ready to start the day?");
if ($choice === "y") {
    echo "That's nice! You should make a to-do list!\n";
} else {
    echo "That's alright! Take 15 minutes of self-care or self-love!\n";
}

$choice = ask("Are you feeling better?");
if ($choice === "y") {
    echo "Awesome! You are now ready for a strong day!";
} else {
    echo "I see... Take a test!";
}