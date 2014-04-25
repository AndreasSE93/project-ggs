package clientChat;

import packageManaging.MessageHandler;
import graphicalReference.ChatGUI;

public class ChatShell {
	
	public MessageHandler handler;
	public ChatGUI graphInterface;
	public TreatInput readFromInterface;
	
	public ChatShell() {
		this.handler = new MessageHandler();
		this.graphInterface = new ChatGUI();
		this.readFromInterface = new TreatInput();
	}
	
}
