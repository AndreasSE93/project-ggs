
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.Socket;
import java.net.UnknownHostException;


public class Testclientsocket {
	final String host = "localhost";
	final int portNumber = 8080;
	Socket socket;
	BufferedReader br;
	PrintWriter out;

	
	Testclientsocket() throws IOException{
		try {
			initializeClient();
		} catch (IOException e) {
			throw(e);
			
		}
	
	}

/*
 * Initalizes a socketclient and connects buffers to in and outstreams 
 * throws IOException if connection failed.
 * 
 */
	
	public void initializeClient() throws IOException{	

		socket = new Socket(host, portNumber);
		br = new BufferedReader(new InputStreamReader(socket.getInputStream()));
		out = new PrintWriter(socket.getOutputStream(), true);
		
		
	}
	
	/*
	 * Writes to the outputstream 
	 */
	public void sendMessage (String k){
		
		out.println(k);
		
	}
	
	/*
	 * Waits and receives a message if such exist from the server, throws IOException if server can't answer
	 */
	public String recieveMessage () throws IOException{
		
		String nextLine;
	
		String message = "";
		nextLine = br.readLine();
		if(nextLine == null){
			IOException e = new IOException();
			throw(e);
		}

		message = nextLine;
		return message;
		
		
	}
	
}

