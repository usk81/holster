# holster
command line tool for file management

## Installation


```
go get github.com/usk81/holster
```

OR

```
wget -O- https://github.com/usk81/holster/releases/download/v2.0.0/{os}.tar.gz | tar xz
```

## Getting Started
To initialize.
This command creates directories and a sample hosts file (is named 'default.host') required to use holster.

```
holster init
```

## Usage

If use init command, make default hosts file in ~/.holster/bullet/.
If you need more hosts file groups, you make in this directory and name these files like ***.host

## Commands

NOTICE:
need to use sudo. (or root user)
If you use windows, open command prompt in administrator mode.

### current
show your hosts file. (It is equal to 'sudo cat /etc/hosts')

```
sudo holster current
# 127.0.0.1       localhost
```

### list

```
sudo holster list
# default
```

### push
overwrite your hosts file.

```
sudo holster push (hosts group)
```

e.g.) Overwrite your hosts file to .holster/bullet/default.host
**before**

```
holster current
# 127.0.0.1       localhost
```

**update**

```
sudo holster push default
```

**after**

```
sudo holster current
# ##
# # Host Database
# #
# # localhost is used to configure the loopback interface
# # when the system is booting.  Do not change this entry.
# ##
# 127.0.0.1       localhost
# 255.255.255.255 broadcasthost
# ::1             localhost
```
