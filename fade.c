int LED = 2;
byte rgb[] = { 255, 0, 0};
int interval = 50, dec = 0, inc;

void setup() {
Serial.begin(115200);
}

void printrgb(byte rgb[]) {
Serial.printf("%02x%02x%02x\n", rgb[0], rgb[1], rgb[2]);
delay(interval);
}

void loop() {
inc = (dec + 1) % 3;

while (rgb[inc] < 255){
rgb[inc] = rgb[inc] + 5;
printrgb(rgb);
}

while (rgb[dec] > 0){
rgb[dec] = rgb[dec] - 5;
printrgb(rgb);
}

dec = (dec + 1) % 3;
delay(interval);
}