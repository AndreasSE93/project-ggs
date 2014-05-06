package clientChat;

import javax.swing.JFrame;
import javax.swing.JPanel;


import clientNetworking.NetManager;

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
	
	public void Initiate(JPanel chatPanel, NetManager NM, JFrame j) {
	this.graphInterface.render(chatPanel, NM, j);	
	}
	
	public void ShutDown() {
		
	}
	
}
