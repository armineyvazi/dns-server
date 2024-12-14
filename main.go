package main

import (
	"fmt"
	"github.com/miekg/dns"
	"log"
)

const fixIp = "192.168.1.1"

// handleDNSRequest processes incoming DNS queries.
func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	response := new(dns.Msg)
	response.SetReply(r)

	// Iterate over the questions in the query
	for _, question := range r.Question {
		switch question.Qtype {
		case dns.TypeA: // Handle A record queries
			ip := fixIp // Fixed IP address for demonstration
			rr, err := dns.NewRR(fmt.Sprintf("%s A %s", question.Name, ip))
			if err == nil {
				response.Answer = append(response.Answer, rr)
			} else {
				log.Printf("Failed to create DNS record: %v\n", err)
			}
		default:
			log.Printf("Unsupported query type: %d\n", question.Qtype)
		}
	}

	// Send the response
	err := w.WriteMsg(response)
	if err != nil {
		log.Printf("Failed to write DNS response: %v\n", err)
	}
}

func main() {
	// Create a new DNS server mux
	dnsMux := dns.NewServeMux()
	dnsMux.HandleFunc(".", handleDNSRequest)

	// Define the DNS server
	server := &dns.Server{
		Addr:    ":5321", // Default DNS port
		Net:     "udp",   // UDP for DNS queries
		Handler: dnsMux,
	}

	log.Println("Starting DNS server on :5321...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start DNS server: %v\n", err)
	}
}
