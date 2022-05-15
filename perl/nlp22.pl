use strict;
use warnings;

use utf8;
use File::Spec;

my $filepath = File::Spec->rel2abs('./output/20_britain.txt');
open(FH, '<', $filepath) or die $!;
my @categories;
while (<FH>) {
    if ( $_ =~ /\[\[Category\:(.+)\]\]/ ) {
        my ($category) = split(/\|/, $1);
        push(@categories, $category);
    }
}

foreach my $category (@categories) {
    print "$category\n";
}