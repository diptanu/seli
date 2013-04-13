class yum {
	file {'/etc/yum.repos.d/seli.conf':
		ensure => present,
		source => "puppet:///modules/yum/seli.conf";
	}
}

