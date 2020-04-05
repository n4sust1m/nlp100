#!/usr/bin/perl

use utf8;
binmode STDOUT, ":utf8";

my $str = "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can.";
my @sanitized_str = split(/\W+/, $str);

my @initial_index = (1, 5, 6, 7, 8, 9, 15, 16, 19);
my @sanitized_index = map { $_ - 1 } @initial_index;

my $result = {};
my $i;
for ($i = 0; $i < scalar @sanitized_str; $i++) {
    my $index_r = sprintf("%02d", $i);
    if ( grep { $_ == $i } @sanitized_index ) {
        $result->{$index_r} = substr(@sanitized_str[$i], 0, 1);
    } else {
        $result->{$index_r} = substr(@sanitized_str[$i], 0, 2);
    }
}

for my $i (sort keys %{$result}) {
    print "$i: " . $result->{$i} . "\n";
}