Exec {
  path => ['/usr/bin', '/usr/sbin', '/bin', '/sbin/']
}

node default {
	stage { "first" : before => Stage['main']}
  include truth::enforcer
}



