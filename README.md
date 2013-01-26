Kontl
=====

What the fuck is this?
----------------------

Kontl is a client for [kon.tl](http://kon.tl)'s URL shortening service written in [Go](http://golang.org/) language.

How to use this thing?
----------------------

Compile:

	$ make
	rm -rf *.o *.a *.[568vq] [568vq].out kontl
	6g -o _go_.6 kontl.go 
	6l -o kontl _go_.6

Run:

	$ kontl 
	Usage: kontl [URL]
	
	$ kontl http://jim.geovedi.com
	 Long URL: http://jim.geovedi.com
	Short URL: http://kon.tl/1689


License
-------

I'm tired of this kind of bullshit. I, hereby, choose WTFPL (Do What The Fuck You Want To Public License) for this code.

	           DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
	                   Version 2, December 2004
		
	Everyone is permitted to copy and distribute verbatim or modified
	copies of this license document, and changing it is allowed as long
	as the name is changed.
	
	           DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
	  TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
	
	 0. You just DO WHAT THE FUCK YOU WANT TO.

