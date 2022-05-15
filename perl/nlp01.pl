#!/usr/bin/perl

use utf8;
binmode STDOUT, ":utf8";

my $str = "パタトクカシーー";
my @chars = split(//, $str);        # required array context
print join('', @chars[0, 2, 4, 6]);