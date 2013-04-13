class truth::enforcer {
	if has_role("dev"){
	  notice("I am a dev vm")
		include tools::build
	}
}

