import java.awt.BorderLayout;
import java.awt.Color;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.io.IOException;

import javax.swing.*;

import org.json.JSONException;


public class ChatInterface implements ActionListener {

	JTextField field;
	
	Testclientsocket client; 
	JTextPane chatPanel;
	String message ="";
	boolean reconnect = true;
	String userName;
	
	public static void main(String[] args) {
	
		ChatInterface ci = new ChatInterface();
		ci.createChatGUI();
		
	}
	
	
	/*
	 * Creates a simple chat GUI
	 */
	public void createChatGUI() {
				
		JFrame chat = new JFrame();
		chat.setLayout(new BorderLayout());
		chat.getContentPane().setBackground(Color.LIGHT_GRAY);
		chat.setSize(800,600);
        chat.setLocationRelativeTo(null);
        chat.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        chat.setVisible(true);
        
        userName = (String)JOptionPane.showInputDialog("Write username!");
        
        field = new JTextField(20);
        
        
        chatPanel = new JTextPane();
        chatPanel.setVisible(true);
        chatPanel.setEditable(false);
        chatPanel.setForeground(Color.BLUE);
        chat.add(chatPanel, "North");
        chat.add(field, "South");
        
        testConnection();
		
        
		field.addActionListener(this);
		updateChat();

	}
	/*
	 * Updates the textChat by receiving message from the server, Handles if a connection failed
	 */
	public void updateChat() {
    	while(true){
    	 		try {
				Message chat = client.recieveMessage();
				setText(chat);
			} catch (IOException e) {
				testConnection();
			}
       		
    		
    	}
		
	}
	
	public void setText(Message chat) {
		message = message + "\n[" + new java.util.Date() + "]" + chat.username + ">" + chat.message;
		
		chatPanel.setText(message);
	}
	
	
	/*
	 * Test the connection, if failed lets the user try to connect again 
	 *
	 */
	public void testConnection() {
        while(true){
		try {
			client = new Testclientsocket();
			return;
		} catch (IOException e) {
			int reply = JOptionPane.showConfirmDialog(null,"Server is down! Would you like to reconnect?", "Warning", JOptionPane.YES_NO_OPTION); 
			if (reply == JOptionPane.NO_OPTION)
				System.exit(1);
		}
			
		}
		
		
	}
	
	public void createMessage(String message) {
		Message chat = new Message(message, userName, "Red");
		try {
			client.encodeJSon(chat);
		} catch (JSONException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		
	}

	
	@Override
	public void actionPerformed(ActionEvent arg0) {
		String text = field.getText();
		createMessage(text);
		
		// TODO Auto-generated method stub
		
	}



	
	
}
