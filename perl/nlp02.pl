#!/usr/bin/perl

use utf8;
binmode STDOUT, ":utf8";

my @str1 = split(//, "パトカー");
my @str2 = split(//, "タクシー");

my $length = @str1;

my $i;
my $result = '';
for (my $i = 0; $i < $length; $i++) {
    $result = $result . $str1[$i] . $str2[$i]
}

print $result;
