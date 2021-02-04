package main

public class JavaSSLServer {

static final int port = 8000;

public static void main(String[] args) {


SSLServerSocketFactory sslServerSocketFactory =
(SSLServerSocketFactory)SSLServerSocketFactory.getDefault();

try {
ServerSocket sslServerSocket =
sslServerSocketFactory.createServerSocket(port);
System.out.println("SSL ServerSocket started");
System.out.println(sslServerSocket.toString());

Socket socket = sslServerSocket.accept();
System.out.println("ServerSocket accepted");

PrintWriter out = new PrintWriter(socket.getOutputStream(), true);
try (BufferedReader bufferedReader =
new BufferedReader(
new InputStreamReader(socket.getInputStream()))) {
String line;
while((line = bufferedReader.readLine()) != null){
System.out.println(line);
out.println(line);
}
}
System.out.println("Closed");

} catch (IOException ex) {
Logger.getLogger(JavaSSLServer.class.getName())
.log(Level.SEVERE, null, ex);
}
}

}