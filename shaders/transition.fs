#version 330

// Input vertex attributes (from vertex shader)
in vec2 fragTexCoord;

// Input uniform values
uniform sampler2D texture0;
uniform sampler2D texture1;
uniform float t;

// Output fragment color
out vec4 finalColor;

void main() {
  // Get the pixel color from both textures
  vec4 fromCol = texture(texture0, fragTexCoord);
  vec4 toCol = texture(texture1, fragTexCoord);

  // Perform linear interpolation between the two colors
  finalColor = mix(fromCol, toCol, t);
}