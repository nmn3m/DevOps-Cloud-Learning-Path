# System Administration

## Note 
	man is your best friend :).

### Directory Structure

home directory → all users. 

etc directory → all configuration of my system. 

#### Note
       Everything is a file in linux. 

dev directory → for devices (kernal or os need to access any thing from this directory)

tmp directory → for temporarily file.

usr directory → all shared file between users ( fonts, colors, background and ...etc ).

bin directory → all users binary its called ( for our command like: cat, date etc).

sbin directory → super user bin ( for root command only ).

var directory  → variable files everything is changeable stored in it like log files.

root directory → like home but for root. 

proc directory → all the data of kernel collected from your device (like the ram free space - hardware info and all this stuff  ) like virtual file system, kernal store it about running application.

opt directory → free folder some vendor want to install their file in it to be easy to remove it.

srv directory → its free also like opt in the future if you need to install service you should do in it.

media directory → if you need to access flash or any drive by mounting it.

mnt directory → for floppy disk. 

boot directory → put in it the bootloader files and the kernal.

---
 
### Basic Definition and Commands

tty → tele type terminal  alt + ctrl.

chvt n → change virtual terminal  number.

id  → user id.

date → show date. 

cal → the calendar of the month you can control it.

su - → change user. 

Tree → list the whole directory files in such a tree.

mkdir → make directory (folder in windows).

cp dir or file → copy file or directory from one place to another → cp ~/Desktop/file0.txt  ~/Documents/file1.txt.

mv dir or dir  → move file or directory from one place to another → mv ~/Desktop/file0.txt  ~/Documents/file1.txt.

rm   → delete files. → rm  file0.txt.

rmdir → delete directory → rmdir -rf directory0. 

rm -rf filesOrFolders →   to force to delete.

mv → using mv to rename files and directories.

#### Note 
     Take care from this  rm -rf /  → you can delete the entire system :).

ls -R  → list all directories and subdirectories.

#### Moving easily in terminal.

	Ctrl + a → go to the beginning of command line.
	  
	Ctrl + e → go to the end of command line.
	  
	Ctrl + k → delete from cursor to the end of line.
	 
	Ctrl + u → delete from cursor to the begining of line. 

 
Ctrl + d → logout.
  
Ctrl + l → clear screen terminal.
   
CTRL + Z → stop the process temporarily.
   
#### Poweroff and Reboot
	   
	Shutdown -r now →  reboot
	  
	Shutdown -h now → halt down
	  
	Reboot 
	 
	Systemctrl reboot
	  
	Systemctrl poweroff

  
using semi colon ; → to do two commands in the same time. 

---
