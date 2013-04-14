class yum {
	file {'/etc/yum.repos.d/seli.repo':
		ensure => present,
		source => "puppet:///modules/yum/seli.repo";
	}
}

