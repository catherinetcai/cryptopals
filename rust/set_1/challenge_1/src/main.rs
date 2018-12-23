extern crate base64;
extern crate hex;

static INPUT: &str = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d";

fn main() {
    // Decode hex string to raw bytes. hex::decode returns a Result type.
    let input_bytes = match hex::decode(INPUT) {
    };

    // Put decoded bytes into base64
    let input_base64 = base64::encode(input_bytes);

    println!("{}", input_base64);
}
