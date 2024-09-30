#version 330

// Input vertex attributes (from vertex shader)
in vec2 fragTexCoord;
in vec4 fragColor;

// Input uniform values
uniform sampler2D from;
uniform sampler2D to;
uniform float t;

// Output fragment color
out vec4 finalColor;

void main() {
  // Get the pixel color from both textures
  vec4 fromCol = texture(from, fragTexCoord);
  vec4 toCol = texture(to, fragTexCoord);

  // Perform linear interpolation between the two colors
  finalColor = mix(fromCol, toCol, t);
}