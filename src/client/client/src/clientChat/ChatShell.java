package clientChat;

import packageManaging.MessageHandler;
import graphicalReference.ChatGUI;
import clientNetworking.Connection;

public class ChatShell {
	
	private Connection conn;
	public MessageHandler handler;
	public ChatGUI graphInterface;
	public TreatInput readFromInterface;
	
	public ChatShell(Connection conn) {
		this.conn = conn;
		this.handler = new MessageHandler();
		this.graphInterface = new ChatGUI();
		this.readFromInterface = new TreatInput();
	}
	
}
