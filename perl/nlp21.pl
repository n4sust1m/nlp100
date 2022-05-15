use strict;
use warnings;

use utf8;
use File::Spec;

my $filepath = File::Spec->rel2abs('./output/20_britain.txt');
open(FH, '<', $filepath) or die $!;
my @categories;
while (<FH>) {
    push(@categories, $_) if ( $_ =~ /\[\[Category\:(.+)\]\]/ );
}

foreach my $category (@categories) {
    print "$category\n";
}