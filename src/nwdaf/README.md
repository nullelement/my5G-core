# NWDAF

Initial version of NWDAF (Network Data Analytics Info). It is a network function designed to run on a 5G core like my5G-core or free5GC.
## Installing

Install on a 5G core, currently works on my5G-core.

## Running

1. First, run the NRF network function;
2. Then, run the NWDAF network function. 

## Simulating service consumption

Tested on my5G-core.

### Events Subscription

- http://localhost:29599/eventssubscription

### Analyltics Info

- http://localhost:29599/analyticsinfo

## Fastest way to test

- Install my5G-core using the tutorial bellow but change the line:
    - `git checkout master` by `git checkout nwdaf-001`
    - [Tutorial: my5G-core Setup](https://github.com/LABORA-INF-UFG/SBrT2020-Minicurso1/blob/master/docs/installation-dev-env-setup/core-install.md)
    
## References:
- [my5G Initiative](https://my5g.github.io/)
- [my5G-core](https://github.com/my5G/my5G-core)
- [free5GC](https://github.com/free5gc/free5gc)  
- [Understanding 5G core through an open-source implementation](https://github.com/LABORA-INF-UFG/SBrT2020-Minicurso1)
