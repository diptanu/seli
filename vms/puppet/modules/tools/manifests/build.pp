class tools::build {
	package {['rpm-build','createrepo']:
		ensure => present;
	}
}

