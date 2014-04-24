package clientChat;

import packageManaging.MessageHandler;
import graphicalReference.ChatGUI;
import clientNetworking.Connection;

public class ChatShell {
	
	Connection conn;
	MessageHandler handler;
	ChatGUI graphInterface;
	TreatInput readFromInterface;
	
	public ChatShell(Connection conn) {
		this.conn = conn;
		this.handler = new MessageHandler();
		this.graphInterface = new ChatGUI();
		this.readFromInterface = new TreatInput();
	}
	
}
