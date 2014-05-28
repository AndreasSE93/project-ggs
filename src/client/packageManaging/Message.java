package packageManaging;

/*
 * Used by client to send and receive a chatMessage.
 * 
 */


public class Message {
	 public final int id = 100;
	 public String user;
	 public String message;
	 
	 public Message(){
		 
		 
	 }

	public Message(String message, String user){
		this.message = message;
		this.user = user;
		
		
	}

	
	
	
	
	
}
