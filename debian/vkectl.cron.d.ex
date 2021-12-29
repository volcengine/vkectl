#
# Regular cron jobs for the vkectl package
#
0 4	* * *	root	[ -x /usr/bin/vkectl_maintenance ] && /usr/bin/vkectl_maintenance
