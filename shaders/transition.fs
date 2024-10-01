#version 330

// Input vertex attributes (from vertex shader)
in vec2 fragTexCoord;

// Input uniform values
uniform sampler2D texture0;
uniform sampler2D texture1;
uniform float t;                  // 0 -> 1
uniform float transitionType = 1; // 0 -> right out
                                  // 1 -> left out
                                  // 2 -> top out
                                  // 3 -> bottom out

// Output fragment color
out vec4 finalColor;

vec2 fromTextCoord;
vec2 toTextCoord;

void main() {
  if (abs(transitionType) < 0.1) {
    fromTextCoord = vec2(t + fragTexCoord.x, fragTexCoord.y);
    toTextCoord = vec2(-1 + t + fragTexCoord.x, fragTexCoord.y);
  } else if (abs(transitionType - 1) < 0.1) {
    fromTextCoord = vec2(-t + fragTexCoord.x, fragTexCoord.y);
    toTextCoord = vec2(1 - t + fragTexCoord.x, fragTexCoord.y);
  }

  vec4 fromCol = texture(texture0, fromTextCoord);
  if (fromTextCoord.x < 0.0 || fromTextCoord.x > 1.0) {
    fromCol = vec4(0.0, 0.0, 0.0, 0.0);
  }
  vec4 toCol = texture(texture1, toTextCoord);
  if (toTextCoord.x < 0.0 || toTextCoord.x > 1.0) {
    toCol = vec4(0.0, 0.0, 0.0, 0.0);
  }

  // Perform linear interpolation between the two colors
  finalColor = fromCol + toCol;
}