# The construct

All of the following assumes an already existing construct. It also assumes a 3rd party VPS provider that is trusted to not snoop on the VPS's memory or registers. An HSM could also be used

To create a new construct, an existing construct does the following:
1. Uses the trusted provider to programatically launch a virtual private server
2. Stores new server's ssh credentials in the existing construct's secrets
3. SSHs in and spins up a new construct
4. SCP the construct key to the new construct. The construct could store the key in an HSM.
5. The construct now has access to all the encrypted secrets

PS, I obviously am a noob w/ crypto. This is just a general outline.

## BitLaunch

Bitlaunch allows a user to provision and pay for a VPS programatically.
