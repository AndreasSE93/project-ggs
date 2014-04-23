
import java.awt.Color;
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.Socket;


import org.json.JSONException;
import org.json.JSONObject;



public class Testclientsocket {
	final String host = "130.243.198.228";
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
	public Message recieveMessage () throws IOException{
		
		String nextLine = "";
	
		String message = "";
		int c;
		while((c = br.read()) != '\n'){
			if(c != 0){
				nextLine = nextLine + (char)c;
			}
		}

		message = nextLine;
		
		try {
			return decodeJson(message);
		} catch (JSONException e) {
			return recieveMessage();
			
		} 
		
	}
	
	/* 
	 * Bör kanske abstraheras
	 * 
	 */
	public void encodeJSon(Message mess) throws JSONException{
	    JSONObject obj = new JSONObject();
	    obj.put("message", mess.message);
	    obj.put("color", mess.color);
	    obj.put("username", mess.username);
	    
		String message = obj.toString();
		sendMessage(message);
		
	}
	
	public Message decodeJson(String Encodedmessage) throws JSONException{
		
		JSONObject obj = new JSONObject(Encodedmessage);
		
								
		Message chatMessage = new Message();
		chatMessage.message = obj.getString("message");
		chatMessage.color =  obj.getString("color");
		chatMessage.username = obj.getString("username");
		
		
		return chatMessage;
	}
	
	
}

