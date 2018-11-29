// Convert hex to base64
extern crate base64;
extern crate hex;
use hex::{decode};
use base64::{encode};

fn main() {
    let input = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d";
    // let expected = String::from("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t");
    let encoded = hex_to_base64(&input);
    println!("{}", encoded);
}

fn hex_to_base64(hex_input: &str) -> String {
    let decoded = decode(hex_input.as_bytes());
    encode(decoded)
}
