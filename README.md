# Cube: a toy orchestrator

I recently read the book [building an orchestrator in
go](https://www.manning.com/books/build-an-orchestrator-in-go-from-scratch).
While working on the code I found a few issues (starting in chapter 8) that
caused compile errors. The issues continued for the next chapters. The code 
in this repo fixes the issues. 

In addition, the orchestrator that we built in the book does not come with any
ingress solution and I wanted to be able to map external traffic to my
containers. Here I use Caddy as a load balancer to send http traffic in a round
robin mode. Caddy also does a  periodic health check to only send traffic to
containers that are healthy.

The connectivity between the components of the orchestrator happens via
Tailscale. In the example that I show, the manager runs in the local machine,
one of the workers runs in AWS and the other on prom. Caddy runs in digital
ocean. The machines connect via the Wireguard direct connections Tailscale
provides. 
