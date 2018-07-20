Try to figure out controls for the chromecast via golang. 
Then we can set up a remote server with a custom alexa skill that calls it. 

notes on running chromecast actions via golang. 

- no default supported API for chromecast except for the chromecast play button, ios device, and android device. Have to use a different language (nodecast, pychromecast, and go-cast are available).
- use https://github.com/ninjasphere/go-castv2/ for now. 
- chromecasts have a concept called namespaces (need to figure out what that is)
- chromecasts need to first be discovered. Then once discovered you can send requests to them. device discover runs via multicast DNS (local DNS service in intranet, may not work in private/cloud settings)
- interactions are via tcp, and requests are composed of specified formats defined in the chromecast sdk(though we can't use the sdk, the formats are still consistent).
- to play, you would use a mediacontroller abstraction object and use its commands. These controller objects have channels with the devices for communication. 

- hmmm, so the golang cast thing doesn't permit that i invoke requests directly through the channel because it is, you know, private. May need to modify it so that we can send more requests to it. 


- chromecast sdk requires that you sign up. 