class tools::build {
	package {['rpm-build','createrepo', 'go']:
		ensure => present;
	}
}

