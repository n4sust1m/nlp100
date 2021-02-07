use strict;
use warnings;

use utf8;
use File::Spec;

my $filepath = File::Spec->rel2abs('./output/20_britain.txt');
open(FH, '<', $filepath) or die $!;
my @title_and_hierarchy;
while (<FH>) {
    if ( $_ =~ /^(=+)([^=]+)(=+)/ ) {
        push @title_and_hierarchy, +{
            title => $2,
            hierarchy => length $1,
        };
    }
}

foreach my $target (@title_and_hierarchy) {
    print 'title: ' . $target->{title} . ', ';
    print 'hierarchy: ' . $target->{hierarchy} . ', ';
    print "\n";
}