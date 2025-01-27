# Notice
This repository is a fork of https://github.com/Diegopyl1209/torrentserver-aniyomi but it looks like it is actually a fork of https://github.com/YouROK/TorrServer/tree/master with an android build. Due to this the GPL-3.0 license applies, even if no licence on the code is given.

The main difference this fork aims to achieve, is to remove some calls to Exit(0) or similar. This is mainly to prevent sigsegv for a better native android experience. The main cause could be a used port, so this also allows us to randomize the port for the server. 