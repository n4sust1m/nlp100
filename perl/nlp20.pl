use strict;
use warnings;

use File::Spec;
use JSON;
use utf8;

# read textfile
my $filepath = File::Spec->rel2abs('./assets/jawiki-country.json');
open(FH, '<', $filepath) or die $!;
my @wiki_country;
while (<FH>) {
    push @wiki_country, decode_json $_;
}

my @britain = grep {
    $_->{title} eq 'イギリス'
} @wiki_country;
my $text_britain = $britain[0]->{text};

print $text_britain;