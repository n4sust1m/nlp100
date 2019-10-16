#!/usr/bin/perl
use utf8;
use strict;
use warnings;

my @str = split(//, "パタトクカシーー");

my $i = 0;
for my $v (@str) {

  if($i % 2 == 1) {
    print $v;
  }

  $i = $i + 1;
}