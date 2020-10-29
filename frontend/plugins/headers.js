export default function headers() {
  const token = localStorage.getItem("auth._token.local");

  if (token) {
    return {
      headers: {
        Authorization: token,
      }
    };
  } else {
    return {};
  }
}
