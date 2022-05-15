#!/usr/bin/perl

use utf8;
binmode STDOUT, ":utf8";

my $str = "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics.";

@result = ();
foreach my $word (split(/[^a-zA-Z]/, $str)) {
    $count = split(//, $word);
    push(@result, $count);
}

print '[' . join(', ', @result) . ']';